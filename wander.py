#!/usr/bin/python3
import random
import requests
import time
import json

delay=5
url = "https://whereischarlie.org/position"
headers = {'Content-type': 'application/json', 'Accept': 'text/plain'}
lat = 42.4701
lng = -83.0565
while True:
    lat = lat + random.choice([-.0001, 0, .0001])
    lng = lng + random.choice([-.0001, 0, .0001])
    data = {'lat': lat, 'lng': lng}
    r = requests.post(url, data=json.dumps(data), headers=headers)
    time.sleep(delay)
