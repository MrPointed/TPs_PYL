suffer_from(Sebastian,Tos).
suffer_from(Sebastian,Cansancio).
suffer_from(Ariel,Diarrea).
suffer_from(Hugo,Fiebre).
suffer_from(Hugo,Nauseas).

symptom_of(Tos,Gripe).
symptom_of(Cansancio,Gripe).
symptom_of(Fiebre,Hepatitis).
symptom_of(Nauseas,Hepatitis).
symptom_of(Diarrea,Intoxicacion).

cure(Vitaminas,Cansancio).
cure(Aspirinas,Fiebre).
cure(Jarabe,Tos).
cure(Bismuto,Diarrea).
cure(Antihistaminicos, Nauseas).

%RULES

disease(A,B) :- suffer_from(A,C), symptom_of(C,B).
must_take(A,B) :- suffer_from(A,C), cure(B,C).