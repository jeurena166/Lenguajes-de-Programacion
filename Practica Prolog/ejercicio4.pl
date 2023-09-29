%Defina un predicado llamado partir(Lista, Umbral, Menores, Mayores) para dividir una lista respecto un umbral dado, 
%dejando los valores menores a la izquierda y los mayores a la derecha. Por ejemplo, el resultado de partir la lista 
%[2,7,4,8,9,1] respecto al umbral 6 serían las listas [2,4,1] y [7,8,9].

partir([], _, [], []).
partir([X|Xs], Umbral, [X|Menores], Mayores) :-
    X =< Umbral,
    partir(Xs, Umbral, Menores, Mayores).
partir([X|Xs], Umbral, Menores, [X|Mayores]) :-
    X > Umbral,
    partir(Xs, Umbral, Menores, Mayores).

%?- partir([2, 7, 4, 8, 9, 1], 6, Menores, Mayores).
%Menores = [2, 4, 1],
%Mayores = [7, 8, 9].
