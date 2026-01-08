SELECT name,
       language,
       model,
       mode,
       pronoun
  FROM (
           SELECT DISTINCT LOWER(name) AS name,
                           LOWER(language) AS language,
                           LOWER(model) AS model,
                           LOWER(mode) AS mode,
                           LOWER(pronoun) AS pronoun
             FROM verbos
           UNION ALL
           SELECT DISTINCT LOWER(model) AS name,
                           LOWER(language) AS language,
                           LOWER(model) AS model,
                           LOWER("infinitivo") AS mode,
                           LOWER(pronoun) AS pronoun
             FROM verbos
            GROUP BY model,
                     language,
                     model,
                     pronoun
       )

