;Tercer ejercicio evaluado

;Variables
(define conjunto '(a b c d e f g h))

;Se define la funcion subconjunto
;En esta funcion, se busca de manera recursiva
;todos los elementos de la sublista
(define (subconjunto lista_sub lista_con)
  (cond ((null? lista_sub)
         #t)
        ((member (car lista_sub) lista_con)
         (subconjunto (cdr lista_sub) lista_con))
        (else
         #f)
        )
  )