# flutter-cleaner
I have a lot of flutter projects that toke a lot disk space and I wanted to archive them, so I needed to run the `flutter clean` command on every project directory and it was painfully slow and repititive.
It happened that I was learning (Go)[https://go.dev] so I decided to make a small utility cmd that scans every directory on a given path and runs the command `flutter clean` for me.

# Usage
```sh
flutter-cleaner $rootpath
```
