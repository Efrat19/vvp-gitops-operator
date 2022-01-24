import fileinput
import os
import re

DIR = "pkg/appmanager_apis"
MARKER = '//+kubebuilder:object:generate=true'
TIME_REPLACE_FROM = '"time"'
TIME_REPLACE_TO = 'time "k8s.io/apimachinery/pkg/apis/meta/v1"'


def replace_in_file(file_path, search_text, new_text):
    with fileinput.input(file_path, inplace=True) as f:
        for line in f:
            new_line = line.replace(search_text, new_text)
            print(new_line, end='')


def useRegex(file):
    with open(file) as src:
        input = src.read()
        res = re.search("type [a-zA-Z0-9]+ struct \{", input)
        if res:
            return res.group()
        return False


for filename in os.scandir(DIR):
    if filename.is_file():
        if "model" in filename.name:
            replace_in_file(filename.path, TIME_REPLACE_FROM, TIME_REPLACE_TO)
            marker_replace_from = useRegex(filename)
            if marker_replace_from:
                replace_in_file(filename.path, marker_replace_from,
                                f'{MARKER}\n{marker_replace_from}')
