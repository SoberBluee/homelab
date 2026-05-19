import boto3
import json

s3 = boto3.client('s3')

def  print_to_s3_on_change(event, context): 
    print('Running print_to_s3_on_change')
    print(event)

    return { 
        'status': 200,
        'body': json.dumps('Successfuly ran s3')
    }

