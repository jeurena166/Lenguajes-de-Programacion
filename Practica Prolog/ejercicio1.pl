% Definir sumlist(L, S) que es verdadero si S es la suma de los elementos de L

sumlist([], 0).
sumlist([X|Xs], S) :-
    sumlist(Xs, S1),
    S is X + S1.

%?- subconj([1, 2, 3], [1, 2, 3, 4, 5]).
% true

%?- subconj([1, 2, 3], [4, 5, 6]).
% false
