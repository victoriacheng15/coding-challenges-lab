import pytest
from cut.cli import main
from io import StringIO
import sys


def test_cli_file_input(tmp_path, monkeypatch):
    file = tmp_path / "test.txt"
    file.write_text("a,b,c\n1,2,3\n")
    monkeypatch.setattr(sys, "argv", ["cut.py", str(file), "-d,", "-f1,3"])
    main()  # Should print "a,c\n1,3\n"


def test_cli_stdin_input(monkeypatch):
    input_stream = StringIO("a,b,c\n1,2,3\n")
    monkeypatch.setattr(sys, "stdin", input_stream)
    monkeypatch.setattr(sys, "argv", ["cut.py", "-d,", "-f1,3"])
    main()  # Should print "a,c\n1,3\n"
