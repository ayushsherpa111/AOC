from __future__ import annotations


class Directory:
    dir_name: str
    files: [int]
    sub_folder: [Directory]
    previous: Directory

    def __init__(self, dir_name, previous=None):
        self.dir_name = dir_name
        self.files = []
        self.sub_folder = []
        self.previous = previous

    def get_size(self) -> int:
        size = sum(self.files)
        size += sum([sub.get_size() for sub in self.sub_folder])
        return size

    def add_dir(self, dir: Directory):
        self.sub_folder.append(dir)

    def add_file(self, file: int):
        self.files.append(int(file))

    def find_sub_dir(self, folder_name):
        return list(
            filter(lambda folder: folder.dir_name == folder_name, self.sub_folder)
        )[0]


def set_command(command_line: str) -> str:
    return command_line.strip().split(" ")[1]


def print_tree(start_dir: Directory, indent: int):
    print(" " * indent + start_dir.dir_name)
    for dir in start_dir.sub_folder:
        print_tree(dir, indent + 1)

def get_sizes(folders: [Directory], limit: int) -> [Directory]:
    under_dir = []
    for folder in folders:
        size = folder.get_size()
        if size <= limit:
            under_dir.append(size)
        if len(folder.sub_folder):
            under_dir += get_sizes(folder.sub_folder, limit)
    return under_dir


with open("./input.txt") as inputs:
    root = Directory("/")
    current_dir = root
    previous_dir = None
    current_command = ""
    for line in inputs:
        args = line.strip().split(" ")
        if args[0] == "$":
            current_command = set_command(line)
            if current_command == "ls":
                continue
        if current_command == "cd":
            if args[2] == "..":
                current_dir = current_dir.previous
            else:
                temp = current_dir.find_sub_dir(args[2])
                # previous_dir = current_dir
                current_dir = temp
        if current_command == "ls":
            if args[0] == "dir":
                current_dir.add_dir(Directory(args[1], current_dir))
            else:
                current_dir.add_file(args[0])

    print(sum(get_sizes(root.sub_folder, 100000)))
