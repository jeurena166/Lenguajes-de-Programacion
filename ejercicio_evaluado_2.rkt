;Segundo ejercicio evaluado

;Variables
(define lista_1 '(1 3 5))
(define lista_2 '(2 4 6))

;Se define la funcion y retorna el resultado
;de un ordenamiento de menor a mayor de los
;elementos de las dos listas.
(define (merge lista1 lista2)
  (define resultado (sort (append lista1 lista2) <)) 
  (display resultado)
  )


