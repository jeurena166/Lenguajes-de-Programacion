;Primer ejercicio Evaluado

(define (intereses cap i n)
  (cond((equal? n 0)
        cap)
       (else
        (intereses (* cap (+ 1 i)) i (- n 1)))
       )
    )