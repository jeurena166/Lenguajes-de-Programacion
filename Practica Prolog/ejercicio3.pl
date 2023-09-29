%Defina la función aplanar(L,L2). Esta función recibe una lista con múltiples listas anidadas como elementos otra lista que contendría los mismos elementos de manera lineal (sin listas). 

aplanar([], []).
aplanar([X|Xs], L2) :-
    aplanar(X, X1),
    aplanar(Xs, Xs1),
    append(X1, Xs1, L2).
aplanar(X, [X]) :- \+ is_list(X).
aplanar([], []).
