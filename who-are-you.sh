#!/bin/bash
# Получаем имя персонажа с id=70 из JSON и выводим в формате "name"
curl -s https://01yessenov.yu.edu.kz/assets/superhero/all.json | \
jq -r '.[] | select(.id == 70) | "\"\(.name)\""'