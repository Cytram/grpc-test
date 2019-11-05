from concurrent import futures

import cv2
import cv2 as cv
import numpy as np
from matplotlib import pyplot as plt

import grpc
import requests

from proto.process import processor_pb2
from proto.process import processor_pb2_grpc


def ParseImage(img_bytes):
    img = cv2.imdecode(np.frombuffer(img_bytes, np.uint8), -1)
    return img

def DownloadImage(url):
    resp = requests.get(url, stream=True).raw
    return cv2.imdecode(np.asarray(bytearray(resp.read()), np.uint8), -1)

def GreyScale(img):
    img = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)
    ret, th1 = cv2.threshold(img, 127, 255, cv2.THRESH_BINARY)
    img = cv.medianBlur(img, 5)

    th2 = cv2.adaptiveThreshold(img, 255, cv2.ADAPTIVE_THRESH_MEAN_C,
                                cv2.THRESH_BINARY, 11, 2)
    th3 = cv2.adaptiveThreshold(img, 255, cv2.ADAPTIVE_THRESH_GAUSSIAN_C,
                                cv2.THRESH_BINARY, 11, 2)
    return th1

def ScaleImage(img, scale):
    height, width = str(scale).split("x")
    print(height, width)
    return cv2.resize(img, dsize=(int(height), int(width)),
                      interpolation=cv2.INTER_CUBIC)


uri = ""

if uri:
    img = DownloadImage(uri)
