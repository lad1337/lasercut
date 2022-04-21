import json


LINE_LENGHTS = {0: 16, 1: 16, 2: 13, 3: 15, 4: 12}

k = json.load(open("lasercut-converted.json"))

matrix = {}
line = []
row_index = 0


def print_layout_line():
    l = ", ".join(line)
    print(f"    {l}, \\")


print("#define LAYOUT( \\")
for index, key in enumerate(k["keys"]):
    row, col = key["profile"].split(",")
    row = int(row)
    col = int(col)
    id_ = f"k{index:02}"
    if row not in matrix:
        matrix[row] = {}
    m_row = matrix[row]
    if col not in m_row:
        m_row[col] = id_
    matrix[row] = m_row
    line.append(id_)
    if len(line) == LINE_LENGHTS[row_index]:
        print_layout_line()
        line = []
        row_index += 1


print(") \\\n{ \\")

for index in sorted(matrix.keys()):
    row = matrix[index]
    line = ", ".join(row[col] for col in sorted(row))
    print(f"   {{{line}}}, \\ #{index}")
print("}")
