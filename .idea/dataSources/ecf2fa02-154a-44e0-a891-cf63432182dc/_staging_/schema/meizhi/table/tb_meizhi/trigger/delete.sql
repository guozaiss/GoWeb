CREATE TRIGGER `delete`
AFTER DELETE ON `tb_meizhi`
FOR EACH ROW
  BEGIN
delete from  tb_meizhi where id='123';
END