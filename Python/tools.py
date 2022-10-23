# ------------------------------------------------------------------------------
# Some tools to deal with facebook archive json.
# ------------------------------------------------------------------------------

import json
from datetime import datetime
from pathlib import Path

print_message_from_file = 'message_20.json'


def is_file_exists(filePath):
    return Path(filePath).exists()


def calc_total_messages():
    """Calculate total messages from all json files."""
    total_msg = 0
    loop_count = 0
    mList = []

    # Generate file name as list.
    while True:
        loop_count += 1
        fileName = 'message_' + str(loop_count) + '.json'
        if is_file_exists(fileName):
            mList.append(fileName)
        else:
            break

    # Open each file to calculate messages.
    for i in mList:
        with open(i) as f:
            data = json.load(f)
            a = data['messages']
            total_msg += len(a)
            # print(f"{i}: {len(a)}")

    # Print total messages.
    print(f"Total messages count: {total_msg}")


def encode_to_human_raedable(encoded_content):
    """Convert facebook archive string to normal string.

    The below code example will be converted to "想我啦"
    data = r'"\u00e6\u0083\u00b3\u00e6\u0088\u0091\u00e5\u0095\u00a6"' 
    print(json.loads(encoded_content).encode('latin1').decode('utf8'))

    Args:
        encoded_content (json object): Something like message['content']

    Returns:
        string: Human readable string.
    """
    return encoded_content.encode('latin1').decode('utf8')


# Print messages from archieve using following format.
# 1970-01-01 00:00:00 <Sender> Message
def print_all_message_from_single_json_file():
    with open(print_message_from_file) as f:
        data = json.load(f)

        for message in data['messages']:
            # Loop through all messages.

            # Print timestamp
            ts = message['timestamp_ms'] / 1000
            print(datetime.fromtimestamp(ts).strftime('%Y-%m-%d %H:%M:%S'),
                end=' ')

            # Print sender' name
            sender_name = encode_to_human_raedable(message['sender_name'])
            print('<%s>' % sender_name, end=' ')

            # Detect message type than print it.
            if 'content' in message:
                # Text
                print(encode_to_human_raedable(message['content']))
            elif 'photos' in message:
                print(message['photos'][0]['creation_timestamp'],
                    message['photos'][0]['uri'])
            elif 'sticker' in message:
                print(message['sticker'])

if __name__ == "__main__":
    calc_total_messages()
    # print_all_message_from_single_json_file()
