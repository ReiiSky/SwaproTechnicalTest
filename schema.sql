CREATE TABLE
  "department" (
    "department_id" SERIAL NOT NULL,
    "department_name" VARCHAR(255) NOT NULL,
    primary key (department_id)
  );

-- End of Department
CREATE TABLE
  "position" (
    position_id serial NOT NULL,
    department_id integer NOT NULL,
    name character varying(255) NOT NULL,
    primary key (position_id)
  );

create index
  position_name_idx on position(name);

alter table
  position
add
  constraint position_department_id_fk foreign key (department_id) references department (department_id);

-- End Of Position
CREATE TABLE
  "employee" (
    "employee_id" SERIAL NOT NULL,
    "employee_code" VARCHAR(255) NOT NULL,
    "position_id" INTEGER NULL,
    "superior_id" INTEGER NULL,
    "name" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    primary key (employee_id)
  );

create index employee_employee_code_idx on employee(employee_code);

alter table
  employee
add
  constraint employee_position_id_fk foreign key (position_id) references position (position_id);

-- End Of Employee
CREATE TABLE
  "location" (
    "location_id" SERIAL NOT NULL,
    "name" VARCHAR(255) NOT NULL,
		primary key (location_id)
  );


-- End Of Location
CREATE TABLE
  "changelog" (
    "log_name" VARCHAR(255) CHECK (
      "log_name" IN (
        'attendance',
        'employee',
        'location',
        'position',
        'department'
      )
    ) NOT NULL,
    "id" BIGINT NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "created_by" VARCHAR(255) NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL,
    "updated_by" VARCHAR(255) NULL,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL,
    primary key (log_name, id)
  );

-- End Of Changelog
CREATE TABLE
  "attendance" (
    "attendance_id" bigserial NOT NULL,
    "employee_id" INTEGER NOT NULL,
    "location_id" INTEGER NOT NULL,
    "absent_in" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "absent_out" TIMESTAMP(0) WITHOUT TIME ZONE NULL,
    primary key (attendance_id)
  );

alter table
  attendance
add
  constraint attendance_employee_id_fk foreign key (employee_id) references employee (employee_id);

alter table
  attendance
add
  constraint attendance_location_id_fk foreign key (location_id) references location (location_id);