import time
from os import system
from watchdog.observers import Observer
from watchdog.events import FileSystemEventHandler


class MyHandler(FileSystemEventHandler):
    def __init__(self, ignore_list):
        self.ignore_list = ignore_list

    def on_modified(self, event):
        if event.is_directory:
            return  # Ignore non-file events

        print(f"File {event.src_path} has been modified.")

        # Check if the modified file is in the ignore list
        for ignore_string in self.ignore_list:
            if ignore_string in event.src_path:
                print("Ignoring file")
                return
            
        system("npm run build")
        print("Build finished")


if __name__ == "__main__":
    # Define your list of files to ignore
    ignore_list = ["vite.config.ts", "dist"]

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