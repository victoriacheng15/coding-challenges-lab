from setuptools import setup, find_packages

setup(
    name="cut-command",
    version="0.1",
    packages=find_packages(),
    entry_points={
        "console_scripts": [
            " pycut=cut.cli:main",
        ],
    },
)
