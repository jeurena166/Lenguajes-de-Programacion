;Cuarto ejercicio evaluado

(define lista_elementos '(1 2 3 4 5 6 7 8 9))

(define (eliminar_elemento e l)
  (map (lambda (n)
         (if (equal? n e) '() n))
       l)
  )
  