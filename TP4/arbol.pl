
%PROPOSITIONS/FACTS

is_parent(Luis,Carolina).
is_parent(Luis,Jimena).

is_parent(Laura,Carolina).
is_parent(Laura,Jimena).

is_parent(Lalo,Sebastian).
is_parent(Patricia,Sebastian).

is_parent(Sebastian,Alejo).
is_parent(Sebastian,Lautaro).
is_parent(Sebastian,Victoria).
is_parent(Sebastian,Constanza).

is_parent(Carolina,Alejo).
is_parent(Carolina,Lautaro).
is_parent(Carolina,Victoria).
is_parent(Carolina,Constanza).

is_parent(Martin,Ignacio).
is_parent(Martin,Fiorella).

is_parent(Jimena,Ignacio).
is_parent(Jimena,Fiorella).

is_parent(Luis,Carolina).
is_parent(Luis,Jimena).

is_woman(Carolina).
is_woman(Jimena).
is_woman(Fiorella).
is_woman(Constanza).
is_woman(Victoria).
is_woman(Patricia).
is_woman(Laura).

%DA RULES Hehe

is_mother(A,B) :- is_parent(A,B), is_woman(A).
is_daughter(B,A) :- is_parent(A,B), is_woman(B).
is_son(B,A) :- is_parent(A,B), \+ is_woman(A).
is_brother(B,C) :- is_parent(A,B),is_parent(A,C).
is_uncle(B,C) :- is_parent(A,B), is_brother(A,C).
is_grandson(B,C) :- is_parent(A,B), is_parent(C,A).
is_grandparent(A,B) :- is_grandson(B,A).