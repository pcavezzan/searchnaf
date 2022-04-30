searchnaf
========

![example workflow](https://github.com/pcavezzan/searchnaf/actions/workflows/build.yml/badge.svg)

Petit utilitaire permettant de récupérer le code naf (activité principal) associée à l'établissement (SIRET) ciblé.

## Usage


### A l'aide un SIRET

```shell
./searchnaf --siret=38432646800343         
siret;naf
38432646800343;71.12B
```

### A l'aide d'un fichier

```shell
./sirenctl -f ./example/input.csv -o ./output/result.csv
```
Un fichier `result.csv` est alors créé dans le répertoire `output`:
```shell
cat output/result.csv 
siret;naf
38432646800343;71.12B
75088177300014;71.12B
75088177300022;71.12B
75149298400016;96.04Z
```

