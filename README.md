This is a tool for fuzzing XSS vulnerabilities based on genetic algorithm.
It's named after Darlene Alderson from Mr. Robot TV series.

## Usage

Basic:
```plain
darlene -url "http://127.0.0.1:8384/reflected2.php?url="
```

With proxy:
```plain
darlene -url "http://127.0.0.1:8384/reflected2.php?url=" -proxy "http://127.0.0.1:8080"
```