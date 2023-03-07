#!/usr/bin/env python3

import json
import os
import subprocess
from datetime import datetime

import boto3
from sense_hat import SenseHat

out_dir = 'output'

# change cwd to script location
os.chdir(os.path.dirname(os.path.abspath(__file__)))

# load config
with open("config.json") as config_file:
    config = json.load(config_file)


def log(msg):
    print(f'sense_hat_collector: {msg}')


def c_to_f(c):
    return c*9.0/5.0 + 32


def dump_sense_hat_data():
    log('getting sense hat data')
    sense = SenseHat()
    sense.clear()
    return {
        'timestamp': datetime.utcnow().isoformat() + 'Z',
        'provider': 'Sense Hat',
        'humidity': sense.get_humidity(),  # percent
        'pressure': sense.get_pressure(),  # millibars
        'temperature': c_to_f(sense.get_temperature()),
        'temperature_from_humidity': c_to_f(sense.get_temperature_from_humidity()),
        'temperature_from_pressure': c_to_f(sense.get_temperature_from_pressure()),
    }


def write_sense_data_data():
    data = dump_sense_hat_data()
    file_name = f'{out_dir}/data-{datetime.utcnow().timestamp()}.json'
    log(f'Writing data to {file_name}')
    with open(file_name, 'w') as outfile:
        json.dump(data, outfile, indent=2)


def upload_file_sqs(file_name, sqs):
    log(f'Uploading {file_name} to sqs')
    with open(file_name, 'r') as infile:
        sqs.send_message(
            QueueUrl=config["sqsUrl"],
            MessageBody=infile.read()
        )
    log(f'Deleting {file_name}')
    os.remove(file_name)


def sync_sqs():
    sqs = boto3.client('sqs')
    for file_name in os.listdir(out_dir):
        upload_file_sqs(f'{out_dir}/{file_name}', sqs)

def sync_local():
    p = subprocess.Popen("./sense-hat-sync")
    p.wait()

def main():
    write_sense_data_data()
    # sync_sqs()
    sync_local()


main()
