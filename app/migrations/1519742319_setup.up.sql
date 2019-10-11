

-- DROP TABLE IF EXISTS public.income CASCADE;
CREATE TABLE public.incomes(
	id serial NOT NULL,
    description varchar(30) NOT NULL,
    amount numeric(15,2) NOT NULL,
    time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
	CONSTRAINT pk_income_id PRIMARY KEY (id)

);

-- DROP TABLE IF EXISTS public.income CASCADE;
CREATE TABLE public.expenses(
	id serial NOT NULL,
    description varchar(30) NOT NULL,
    amount numeric(15,2) NOT NULL,
    time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
	CONSTRAINT pk_expense_id PRIMARY KEY (id)

);

