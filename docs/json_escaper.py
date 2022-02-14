"""
$ pip install pyperclip
# ex) python json_escaper.py -f plantuml.json
$ python json_escaport.py -f <파일명>
This scripts read "filename" and convert file to one string with escape chars.
And finally copies text to your clipboard by pyperclip.
"""
import argparse
import pathlib

import pyperclip  # noqa

parser = argparse.ArgumentParser()
parser.add_argument("-f", type=pathlib.Path, help="Target file")
args = parser.parse_args()


def get_string_with_escape(my_string):
    """
    with linebreak(\n)
    ' '*4 -> \t
    ' '*3 -> \t
    ' '*2 -> \t
    ' ' -> do nothing
    """
    # string_with_linebreak = repr(my_string)
    string_with_linebreak_and_without_quote = repr(my_string).replace("'", "")
    return (
        string_with_linebreak_and_without_quote.replace(" " * 4, r"\t").replace(" " * 3, r"\t").replace(" " * 2, r"\t")
    )


if __name__ == "__main__":
    with args.f.open("r", encoding="utf-8") as f:
        old_lines = f.readlines()
    new_lines = "".join(old_lines)
    result_with_escape = get_string_with_escape(new_lines)

    print(f"RESULT = {result_with_escape}")
    pyperclip.copy(result_with_escape)
    pyperclip.paste()
