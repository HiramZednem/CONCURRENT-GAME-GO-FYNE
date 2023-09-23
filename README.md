Mi juego va a ser un tux que se mueve de izquierda a derecha
Y le caen imagenes de windows, si windows le pega a tux -> tux se enoja y el juego se detiene
Va a haber un temporizador y cada temporizador spawneara un windows en posicion aleatoria

Hilos:
Tux
Windows
Cronometro

El hilo del cronometro, cada 10 segundos spawneara un hilo de windows, el hilo de windows no va a estar en un ciclo for,
simplemente su tarea va ser desplazarse hacia abajo, y cuando llegue hasta abajo el hilo muere

El hilo del cronometro va a ser mostrar el cronometro y cada 8 segundos spawnear un windows

Tux:
PosX, PosY, status, peel -> MoveRight, MoveLeft, ChangeLayout

Windows:
PosX, PosY, speed, peel -> Collisioned ( tux.stop, cronometro.stop )

Cronometro:
time -> NewWindows -> stop, start