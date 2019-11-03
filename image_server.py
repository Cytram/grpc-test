from concurrent import futures
import logging

import cv2
import numpy as np
from matplotlib import pyplot as plt

import grpc

from proto.process import processor_pb2
from proto.process import processor_pb2_grpc

class ImageScaler(processor_pb2_grpc.ProcessorServicer):

    def __ParseImage(self, img_bytes):
        img = cv2.imdecode(np.frombuffer(img_bytes, np.uint8), -1)
        return img

    def __GreyScale(self, img):
        img = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)
        return img

    def ProcessImage(self, request, context):
        for key, value in context.invocation_metadata():
            print('Received initial metadata: key=%s value=%s' % (key, value))
        print("")

        img = self.__ParseImage(request.image.content)

        height, width, channels = img.shape
        print(height, width, channels)

        res = cv2.resize(img, dsize=(1024, 768), interpolation=cv2.INTER_CUBIC)

        img = self.__GreyScale(res)
        img_str = cv2.imencode('.jpg', res)[1].tostring()
        
        return processor_pb2.ProcessImageReply(content=bytes(img_str))


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    processor_pb2_grpc.add_ProcessorServicer_to_server(ImageScaler(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
