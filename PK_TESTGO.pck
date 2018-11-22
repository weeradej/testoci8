create or replace package PK_TESTGO is

  -- Author  : WEERADEJ
  -- Created : 21/11/2018 12:24:18
  -- Purpose : Test golang call procedure

  PROCEDURE GETPRODUCTS(retval OUT VARCHAR2);
end PK_TESTGO;
/
create or replace package body PK_TESTGO is

  PROCEDURE GETPRODUCTS(retval OUT VARCHAR2) is
  begin
    select 'Hello world! ' into retval from dual;
  end GETPRODUCTS;
end PK_TESTGO;
/
