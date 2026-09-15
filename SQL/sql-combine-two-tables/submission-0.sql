-- Write your query below
select first_name, last_name, city, state from person p  left join address add on 
p.person_id = add.person_id