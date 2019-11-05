#!/usr/bin/python3

import socket
import struct

url="35.239.96.98"
port=5000
ai = socket.getaddrinfo(url, port)
s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
data = bytearray(b'm\xe1)B\xf7\x1c\xa6\xc2')
s.sendto(data, ai[0][-1])


