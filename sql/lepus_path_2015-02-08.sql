
update `admin_menu` set `menu_url` = CONCAT('lp_', `menu_url`) WHERE `menu_url` LIKE 'mongodb%';
update `admin_menu` set `menu_url` = CONCAT('lp_', `menu_url`) WHERE `menu_url` LIKE 'oracle%';
update `admin_menu` set `menu_url` = CONCAT('lp_', `menu_url`) WHERE `menu_url` LIKE 'os%';
update `admin_menu` set `menu_url` = CONCAT('lp_', `menu_url`) WHERE `menu_url` LIKE 'mysql%';
update `admin_menu` set `menu_url` = CONCAT('lp_', `menu_url`) WHERE `menu_url` LIKE 'redis%';


update `admin_privilege` set `action` = CONCAT('lp_', `action`) WHERE `action` LIKE 'mongodb%';
update `admin_privilege` set `action` = CONCAT('lp_', `action`) WHERE `action` LIKE 'oracle%';
update `admin_privilege` set `action` = CONCAT('lp_', `action`) WHERE `action` LIKE 'os%';
update `admin_privilege` set `action` = CONCAT('lp_', `action`) WHERE `action` LIKE 'mysql%';
update `admin_privilege` set `action` = CONCAT('lp_', `action`) WHERE `action` LIKE 'redis%';

