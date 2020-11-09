import re


PACKAGE_NAME = "colorname"
OUT_FILE = 'colortable.go'
LINE_RE = re.compile('(.+)\s+#([\dA-Fa-f]{6})')
DATA_FILES = [
    './data/a-f.txt',
    './data/g-m.txt',
    './data/n-z.txt'
]


def main():
    data = []
    for data_file in DATA_FILES:
        with open(data_file, 'r', encoding='utf-8') as f:
            data.extend(f.readlines())

    data = [map_line(l) for l in data]
    data = [to_go_map_line(l, 1) for l in data if len(l) > 0]

    out = """package {}

var colorNames = map[string]int {{
{}
}}
    """.format(PACKAGE_NAME, '\n'.join(data))

    with open(OUT_FILE, 'w', encoding='utf-8') as f:
        f.write(out)


def map_line(line):
    match = LINE_RE.match(line)
    if not match:
        print("Failed matching line: {}".format(line))
        return ()
    return match.groups()


def to_go_map_line(line, tab_count):
    name, hex_val = line
    return '{}"{}": 0x{},'.format('\t' * tab_count, name, hex_val)


if __name__ == '__main__':
    main()
