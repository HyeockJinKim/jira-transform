import sys
from jira2markdown import convert

if __name__ == '__main__':
    input_string = "".join(sys.argv[1:])
    print(convert(input_string))
