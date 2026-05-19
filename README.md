# gendiff

### Hexlet tests and linter status
[![Hexlet Check](https://github.com/ibir-bit/go-project-244/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/ibir-bit/go-project-244/actions)

### Go CI (lint, build, tests)
[![Go CI](https://github.com/ibir-bit/go-project-244/actions/workflows/go.yml/badge.svg)](https://github.com/ibir-bit/go-project-244/actions/workflows/go.yml)

### SonarCloud
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ibir-bit_go-project-244&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=ibir-bit_go-project-244)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=ibir-bit_go-project-244&metric=coverage)](https://sonarcloud.io/summary/new_code?id=ibir-bit_go-project-244)

---

## О проекте

`gendiff` — CLI-утилита для сравнения двух JSON-файлов и вывода различий.


Основной функционал:
- Поддержка JSON и YAML форматов
- Рекурсивное сравнение вложенных структур
- Вывод различий в читаемом формате stylish
- Визуальные маркеры изменений (+ добавлено, - удалено)
- Обработка null значений

Технические детали:
- Парсер файлов с автоматическим определением формата
- Построение дерева различий для вложенных объектов
- Форматирование вывода с правильными отступами
- Обработка пустых строк и null значений
- Модульные тесты с покрытием  

Пример использования:

```bash
./bin/gendiff file1.json file2.json