import sys
import requests

DEFAULT_BASE_URL = "http://localhost:8080"

def list_files(base_url):
    response = requests.get(f"{base_url}/files")
    if response.status_code == 200:
        files = response.json()
        if files:
            print("Files:")
            for file in files:
                print(f"- {file}")
        else:
            print("No files uploaded.")
    else:
        print("Failed to list files:", response.status_code, response.text)

def upload_file(base_url, filename):
    with open(filename, 'rb') as f:
        files = {'file': (filename, f)}
        response = requests.post(f"{base_url}/upload", files=files)
    if response.status_code == 200:
        print(f"File {filename} uploaded successfully.")
    else:
        print("Failed to upload file:", response.status_code, response.text)

def delete_file(base_url, filename):
    response = requests.delete(f"{base_url}/files/{filename}")
    if response.status_code == 200:
        print(f"File {filename} deleted successfully.")
    elif response.status_code == 404:
        print(f"File {filename} not found.")
    else:
        print("Failed to delete file:", response.status_code, response.text)

def main():
    if len(sys.argv) < 2:
        print("Usage: wb [-url <base_url>] [list | up -f <filename> | remove <filename>]")
        sys.exit(1)

    base_url = DEFAULT_BASE_URL

    if sys.argv[1] == "-url":
        if len(sys.argv) >= 3:
            base_url = sys.argv[2]
            command_args = sys.argv[3:]
        else:
            print("Usage: wb [-url <base_url>] [list | up -f <filename> | remove <filename>]")
            sys.exit(1)
    else:
        command_args = sys.argv[1:]

    command = command_args[0]

    if command == "list":
        list_files(base_url)
    elif command == "up" and len(command_args) == 3 and command_args[1] == "-f":
        upload_file(base_url, command_args[2])
    elif command == "remove" and len(command_args) == 2:
        delete_file(base_url, command_args[1])
    else:
        print("Invalid command or arguments.")
        print("Usage: wb [-url <base_url>] [list | up -f <filename> | remove <filename>]")

if __name__ == "__main__":
    main()
