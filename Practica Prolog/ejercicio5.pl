%Implemente un predicado que, a partir de una lista de cadenas string, filtre aquellas que contengan una subcadena que el usuario indique en otro argumento. 

sub_cadenas(_, [], []).
sub_cadenas(Subcadena, [Cadena|CadenasRestantes], Filtradas) :-
    sub_string(Cadena, _, _, _, Subcadena),
    sub_cadenas(Subcadena, CadenasRestantes, FiltradasRestantes),
    Filtradas = [Cadena|FiltradasRestantes].
sub_cadenas(Subcadena, [_|CadenasRestantes], Filtradas) :-
    sub_cadenas(Subcadena, CadenasRestantes, Filtradas).

%?- sub_cadenas("la", ["la casa", "el perro", "pintando la cerca"], Filtradas).
%Filtradas = ["la casa", "pintando la cerca"].

%?- sub_cadenas("a", ["manzana", "banana", "pera"], Filtradas).
%Filtradas = ["manzana", "banana", "pera"].
