import time
from os import system
from watchdog.observers import Observer
from watchdog.events import FileSystemEventHandler
from colorama import Back


class MyHandler(FileSystemEventHandler):
    def __init__(self, ignore_list):
        self.ignore_list = ignore_list

    def on_modified(self, event):
        if event.is_directory:
            return  # Ignore non-file events

        print(f"{Back.BLUE}File {Back.MAGENTA}{event.src_path}{Back.BLUE} has been modified.{Back.RESET}")

        # Check if the modified file is in the ignore list
        for ignore_string in self.ignore_list:
            if ignore_string in event.src_path:
                print(Back.YELLOW + "Ignoring file" + Back.RESET)
                return
            
        system("npm run build")
        print(Back.GREEN + "Build finished" + Back.RESET)


if __name__ == "__main__":
    ignore_list = ["vite.config.ts", "dist", "node_modules"]

    path_to_watch = "./"
    event_handler = MyHandler(ignore_list)
    observer = Observer()
    observer.schedule(event_handler, path=path_to_watch, recursive=True)
    observer.start()

    try:
        while True:
            time.sleep(1)
    except KeyboardInterrupt:
        observer.stop()

    observer.join()