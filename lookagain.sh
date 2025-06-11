find . -type f -name "*.sh" | sed 's/\.sh$//' | sed 's/^.*\///' | sort -r
echo
