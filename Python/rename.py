# ------------------------------------------------------------------------------
# Rename photos to readable name from facebook archive.
# ------------------------------------------------------------------------------

import os
import json
from datetime import datetime
from pathlib import Path
from distutils.dir_util import copy_tree

# Rename json file's name here in case u have multiple json files.
json_file_name = 'message_20.json'


def generate_all_files_name_as_list():
    """Generate a list of all json files name in current working directory.
    
    Returns:
        list: list of all json files name in current working directory.
    """
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

    return mList


def backup_original_photo():
    """Backup original photo to photos_original folder."""

    # Copy photos to photos_original folder if it does not exist.
    if not os.path.exists('photos_original'):
        copy_tree("photos", "photos_original")


def find_nth(mString, target, n):
    """Find nth occurrence of string.

    Args:
        mString (str): string to search
        target (str): target string
        n (int): occurrence to find

    Returns:
        int: index of nth occurrence of target in mString
    """
    start = mString.find(target)
    while start >= 0 and n > 1:
        start = mString.find(target, start + len(target))
        n -= 1
    return start


def find_photo_uri(message):
    nth = find_nth(message, '/', 3)
    return message[nth + 1:]


def convert_timestamp_to_datetime(timestamp):
    return datetime.fromtimestamp(timestamp).strftime('%Y%m%d_%H%M%S')


def is_file_exists(filePath):
    return Path(filePath).exists()


def rename_file(filePath, convertedDatetime):
    """Rename file with datetime if the file exists.

    Args:
        filePath (str): Photos uri from JSON file, looks like 
        "photos/file_name_here.jpg".
        datetime (str): The format are YYYYMMDD_HHMMSS.
    """
    my_file = Path(filePath)

    try:
        creation_timestamp = convert_timestamp_to_datetime(convertedDatetime)
        newFileName = "IMG_{}".format(creation_timestamp)
        newFileNameWithPath = 'photos/' + newFileName + my_file.suffix

        # Rename file then print the result.
        os.rename(filePath, newFileNameWithPath)
        print('[OK] {}, {}'.format(newFileName, filePath))

    except FileExistsError:
        # If the file already exists,
        # decreasing the timestamp by 1 second or more.

        secondsToAvoid = 1

        creation_timestamp = convert_timestamp_to_datetime(convertedDatetime -
                                                           secondsToAvoid)
        newFileName = "IMG_{}".format(creation_timestamp)
        newFileNameWithPath = 'photos/' + newFileName + my_file.suffix

        while (is_file_exists(newFileNameWithPath)):
            creation_timestamp = convert_timestamp_to_datetime(
                convertedDatetime - secondsToAvoid)
            newFileName = "IMG_{}".format(creation_timestamp)
            newFileNameWithPath = 'photos/' + newFileName + my_file.suffix
            secondsToAvoid += 1

        # Rename file then print the result.
        os.rename(filePath, newFileNameWithPath)
        print('[Duplicated, OK] {}, {}'.format(newFileName, filePath))


if __name__ == "__main__":
    # Backup the old photos folder before the operation.
    backup_original_photo()

    with open(json_file_name, "r") as f:
        data = json.load(f)

    for message in data['messages']:
        # Loop through all messages.

        if 'photos' in message:
            # If the message type is photo, then rename the file.

            for i in range(0, len(message['photos'])):
                # Loop through photos,
                # sometimes a single message has more than one photo.

                filePath = find_photo_uri(message['photos'][i]['uri'])
                if (is_file_exists(filePath)):
                    rename_file(filePath,
                                message['photos'][i]['creation_timestamp'])