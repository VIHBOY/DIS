# Tarea 1 Sistemas Dsitribuidos


## Comenzando 🚀

_Estas instrucciones te permitirán obtener una copia del proyecto en funcionamiento en tus máquinas virtuales para propósitos de desarrollo y pruebas._

### Pre-requisitos 📋

_Que cosas necesitas para instalar el software y como instalarlas_

```
Las maquinas virtuales ya tienen todo lo necesario para realizar pruebas.
```

### Pasos a Seguir 🔧

_Se debe seguir estos pasos para poder ejecutar la tarea_

```
1. En la Maquina dist25, se debe ejecutar go run logistica.go

Lo que le otorgara a la maquina el rol de LOGISTICA

2. En la Maquina dist28, se debe ejecutar go run finanza.go

Lo que le otorgara a la maquina el rol de FINANCIERO

3. En la Maquina dist26, se debe ejecutar go run cliente.go

Lo que le otorgara a la maquina el rol de CLIENTE

4. En la Maquina dist27, se debe ejecutar go run camion.go

Lo que le otorgara a la maquina el rol de CAMION
```

## Consideraciones Logistica ⚙️

En todo momento se debe mantener encendido el servidor.

## Consideraciones Clientes ⚙️

```
make
java MergeSort
```
_Se le pedira que entregue por terminal una lista con el siguiente formato: 1,3,9,4,2_

_Luego se mostrara por la misma terminal la lista ordenada_

### Explicación de MergeSort con Multihebra
_Al momento de dividir el arreglo con la tecnica merge, durante la aplicacion de mergesort se aplicaran hebras, para que todos los subarreglos se ordenen y junten, ya que el algoritmo original va ordenando por subarreglo, mientras los otros se quedan esperando, lo cual no seria optimo comparado con lo descrito al principio_

## Construido con 🛠️

* [VS code] - Editor de texto

## Autores ✒️

* **Joaquin Concha** - 201773569-4 *VIHBOY*
    -_Problema 1_
* **Renato Bassi** - 201773521-K *bassisi*
    -_Problema 2_  