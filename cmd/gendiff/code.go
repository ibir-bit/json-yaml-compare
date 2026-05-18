package code

import "code/pkg/gendiff"

// GenDiff — это обертка над твоей основной функцией,
// чтобы тесты Хекслета могли найти её в пакете 'code'
var GenDiff = gendiff.GenDiff
