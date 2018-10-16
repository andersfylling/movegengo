# MovegenGo [![Documentation](https://godoc.org/github.com/chessmodule/movegengo?status.svg)](http://godoc.org/github.com/chessmodule/movegengo)

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)[![forthebadge](https://forthebadge.com/images/badges/for-you.svg)](https://forthebadge.com)

## Health
| Branch       | Build status  | Code climate | Go Report Card | Code Coverage |
| ------------ |:-------------:|:---------------:|:-------------:|:----------------:|
| master     | [![CircleCI](https://circleci.com/gh/chessmodule/movegengo/tree/master.svg?style=shield)](https://circleci.com/gh/chessmodule/movegengo/tree/master) | [![Maintainability](https://api.codeclimate.com/v1/badges/f28b832369e4027522e7/maintainability)](https://codeclimate.com/github/chessmodule/movegengo/maintainability) | [![Go Report Card](https://goreportcard.com/badge/github.com/chessmodule/movegengo)](https://goreportcard.com/report/github.com/chessmodule/movegengo) | [![Test Coverage](https://api.codeclimate.com/v1/badges/f28b832369e4027522e7/test_coverage)](https://codeclimate.com/github/chessmodule/movegengo/test_coverage) |

## Bitboard layout
```

# Chess board layout

63	62	61	60	59	58	57	56
55	54	53	52	51	50	49	48
47	46	45	44	43	42	41	40
39	38	37	36	35	34	33	32
31	30	29	28	27	26	25	24
23	22	21	20	19	18	17	16
15	14	13	12	11	10	09	08
07	06	05	04	03	02	01	00


# uint64 layout by index

63	62	61	60	59	58	57	56  ..  07	06	05	04	03	02	01	00

```
The first bit (index `0`) of a bitboard or uint64 is `00`, and the last is `63`.