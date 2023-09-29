% Definir sumlist(L, S) que es verdadero si S es la suma de los elementos de L

sumlist([], 0).
sumlist([X|Xs], S) :-
    sumlist(Xs, S1),
    S is X + S1.

%?- sumlist([], S).
%S = 0.

%?- sumlist([1.5, 2.7, 3.2, 4.1], S).
%S = 11.5.
