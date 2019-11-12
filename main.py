from concurrent import futures

import cv2
import cv2 as cv
import numpy as np
from matplotlib import pyplot as plt

import grpc
import requests

from proto.process import processor_pb2
from proto.process import processor_pb2_grpc


class ImageScaler(processor_pb2_grpc.ProcessorServicer):
    def __ParseImage(self, img_bytes):
        img = cv2.imdecode(np.frombuffer(img_bytes, np.uint8), -1)
        return img

    def __DownloadImage(self, url):
        resp = requests.get(url, stream=True).raw
        return cv2.imdecode(np.asarray(bytearray(resp.read()), np.uint8), -1)

    def __GreyScale(self, img):
        img = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)
        ret, th1 = cv2.threshold(img, 127, 255, cv2.THRESH_BINARY)
        img = cv.medianBlur(img, 5)

        return th1

    def __ScaleImage(self, img, scale):
        height, width = str(scale).split("x")
        print(height, width)
        return cv2.resize(img, dsize=(int(height), int(width)),
                          interpolation=cv2.INTER_CUBIC)

    def ProcessImage(self, request, context):
        if request.image.source.http_uri:
            img = self.__DownloadImage(request.image.source.http_uri)
        else:
            img = self.__ParseImage(request.image.content)

        if "x" in request.scale:
            img = self.__ScaleImage(img, request.scale)
   
        if request.grayscale:
            img = self.__GreyScale(img)
        img_str = cv2.imencode('.jpg', img)[1].tostring()

        return processor_pb2.ProcessImageReply(content=bytes(img_str))


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    processor_pb2_grpc.add_ProcessorServicer_to_server(ImageScaler(), server)
    server.add_insecure_port('[::]:50052')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
