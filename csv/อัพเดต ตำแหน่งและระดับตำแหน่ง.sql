ALTER TABLE t005.job_postion ADD position_id_new varchar(100) NULL COMMENT 'ตำแหน่ง HRIS';
ALTER TABLE t005.job_postion CHANGE position_id_new position_id_new varchar(100) NULL COMMENT 'ตำแหน่ง HRIS' AFTER postion_id;
ALTER TABLE t005.job_applicant ADD postion_id_new varchar(20) NULL COMMENT 'ตำแหน่ง hris';
ALTER TABLE t005.job_applicant CHANGE postion_id_new postion_id_new varchar(20) NULL COMMENT 'ตำแหน่ง hris' AFTER postion_id;
ALTER TABLE t005.job_applicant ADD position_level_new varchar(20) NULL COMMENT 'ระดับตำแหน่ง hris';
ALTER TABLE t005.job_applicant CHANGE position_level_new position_level_new varchar(20) NULL COMMENT 'ระดับตำแหน่ง hris' AFTER postion_id_new;
ALTER TABLE t005.job_applicant ADD dep_id_hris varchar(20) NULL;
ALTER TABLE t005.job_applicant CHANGE dep_id_hris dep_id_hris varchar(20) NULL AFTER section_id;
ALTER TABLE t005.job_applicant ADD section_id_hris varchar(20) NULL;
ALTER TABLE t005.job_applicant CHANGE section_id_hris section_id_hris varchar(20) NULL AFTER dep_id_hris;




UPDATE job_applicant 
INNER JOIN hris_employees ON job_applicant.applicant_id = hris_employees.id
SET 
    job_applicant.postion_id_new = hris_employees.position_id,
    job_applicant.position_level_new = hris_employees.position_level_id;


UPDATE job_applicant 
INNER JOIN hris_employees ON job_applicant.applicant_id = hris_employees.id
SET 
    job_applicant.dep_id_hris = hris_employees.department_id,
    job_applicant.section_id_hris = hris_employees.section_id;






postion_id_new
position_level_new



SELECT postion_id
FROM job_applicant
WHERE status_work=1
GROUP BY postion_id
HAVING COUNT(DISTINCT postion_id_new) > 1;


SELECT postion_id
FROM job_applicant
WHERE 1
GROUP BY postion_id
HAVING COUNT(DISTINCT postion_id_new) > 1;


SELECT postion_id
FROM job_applicant
WHERE status_work=1
GROUP BY postion_id
HAVING COUNT(DISTINCT postion_id_new) = 1;





'p0406','p0408','p0414','p0415','p0416','p0419','p0421','p0422','p0424','p0426','p0427','p0428','p0430','p0432','p0436','p0438','p0439','p0440','p0444','p0446','p0449','p0450','p0452','p0462','p0464','p0467','p0469','p0471','p0472','p0474','p0475','p0478','p0481','p0482','p0489','p0490','p0491','p0493','p0494','p0495','p0496','p0500','p0507','p0511','p0522','p0523','p0524','p0525','p0526','p0530','p0533','p0535','p0536','p0542','p0543','p0544','p0552','p0554','p0557','p0564','p0566','p0569','p0570','p0581','p0582','p0584','p0586','p0587','p0602','p0604','p0606','p0607','p0612','p0613','p0614','p0615','p0619','p0621','p0622','p0623','p0633','p0634','p0635','p0638','p0643','p0652','p0657','p0660','p0662','p0664','p0666','p0670','p0673','p0675','p0682','p0683','p0686','p0687','p0692','p0699','p0701','p0707','p0708','p0709','p0710','p0711','p0713','p0715','p0716','p0722','p0724','p0725','p0726','p0729','p0733','p0735','p0736','p0737','p0743','p0744','p0745','p0746','p0747','p0748','p0750','p0751','p0755','p0758','p0764','p0768','p0769','p0773','p0780','p0782','p0784','p0785','p0787','p0792','p0794','p0795','p0799','p0800','p0801','p0802','p0804','p0805','p0806','p0809','p0811','p0815','p0817','p0821','p0826','p0828','p0833','p0834','p0837','p0839','p0841','p0845','p0846','p0847','p0848','p0851','p0853','p0859','p0860','p0862','p0863','p0873','p0874','p0878','p0882','p0884','p0888','p0890','p0891','p0893','p0894','p0895','p0896','p0897','p0898','p0906','p0909','p0910','p0911','p0913','p0915'

UPDATE job_postion SET job_postion.position_id_new =(SELECT job_applicant.postion_id_new FROM job_applicant WHERE job_applicant.postion_id=job_postion.postion_id Limit 1) WHERE job_postion.postion_id IN ('p0406','p0408','p0414','p0415','p0416','p0419','p0421','p0422','p0424','p0426','p0427','p0428','p0430','p0432','p0436','p0438','p0439','p0440','p0444','p0446','p0449','p0450','p0452','p0462','p0464','p0467','p0469','p0471','p0472','p0474','p0475','p0478','p0481','p0482','p0489','p0490','p0491','p0493','p0494','p0495','p0496','p0500','p0507','p0511','p0522','p0523','p0524','p0525','p0526','p0530','p0533','p0535','p0536','p0542','p0543','p0544','p0552','p0554','p0557','p0564','p0566','p0569','p0570','p0581','p0582','p0584','p0586','p0587','p0602','p0604','p0606','p0607','p0612','p0613','p0614','p0615','p0619','p0621','p0622','p0623','p0633','p0634','p0635','p0638','p0643','p0652','p0657','p0660','p0662','p0664','p0666','p0670','p0673','p0675','p0682','p0683','p0686','p0687','p0692','p0699','p0701','p0707','p0708','p0709','p0710','p0711','p0713','p0715','p0716','p0722','p0724','p0725','p0726','p0729','p0733','p0735','p0736','p0737','p0743','p0744','p0745','p0746','p0747','p0748','p0750','p0751','p0755','p0758','p0764','p0768','p0769','p0773','p0780','p0782','p0784','p0785','p0787','p0792','p0794','p0795','p0799','p0800','p0801','p0802','p0804','p0805','p0806','p0809','p0811','p0815','p0817','p0821','p0826','p0828','p0833','p0834','p0837','p0839','p0841','p0845','p0846','p0847','p0848','p0851','p0853','p0859','p0860','p0862','p0863','p0873','p0874','p0878','p0882','p0884','p0888','p0890','p0891','p0893','p0894','p0895','p0896','p0897','p0898','p0906','p0909','p0910','p0911','p0913','p0915')




UPDATE job_postion SET job_postion.position_id_new =(SELECT hris_positions.id FROM hris_positions WHERE hris_positions.name_th=job_postion.name Limit 1) WHERE job_postion.position_id_new is null


//ข้อมูล config map ตำแหน่ง
SELECT `postion_id`,`name`,`name_eng`,`status`,`position_id_new`,(SELECT hris_positions.name_th FROM hris_positions WHERE hris_positions.id=job_postion.position_id_new) as position_name_th_hris,(SELECT hris_positions.name_en FROM hris_positions WHERE hris_positions.id=job_postion.position_id_new) as position_name_en_hris,(SELECT hris_positions.is_active FROM hris_positions WHERE hris_positions.id=job_postion.position_id_new) as is_active FROM `job_postion` WHERE 1



//ข้อมูลพนักงาน
SELECT `applicant_id`,`th_name`,`th_surname`,`postion_id`,(select job_postion.name FROM job_postion WHERE job_postion.postion_id=job_applicant.postion_id) as postion_id_txt_pis,`postion_id_new`,(select hris_positions.name_th FROM hris_positions WHERE hris_positions.id=job_applicant.postion_id_new) as postion_id_txt_hris,`position_level`,(select node_name_th FROM conf_com_const WHERE config_id=position_level) as position_level_pis,`position_level_new`,(select node_name_th FROM conf_com_const WHERE config_id=position_level_new) as position_level_hris,`status_work` FROM `job_applicant` WHERE 1;



UPDATE job_postion 
INNER JOIN hris_positions ON job_postion.name = hris_positions.name_th and hris_positions.is_active=1
SET job_postion.position_id_new = hris_positions.id;




SELECT 
    ja.`applicant_id` as 'รหัสพนักงาน',
    ja.`th_name` as 'ชื่อ',
    ja.`th_surname` as 'นามสกุล',
    ja.`postion_id` as 'id ตำแหน่ง pis',
    ja.`postion_id_new` as 'id ตำแหน่ง hris',
    ja.`position_level` as 'id ระดับตำแหน่ง pis',
    ja.`position_level_new` as 'id ระดับตำแหน่ง hris',
    ja.dep_id as 'id ฝ่าย pis',
    ja.dep_id_hris as 'id ฝ่าย hris',
    ja.section_id as 'id แผนก pis',
    ja.section_id_hris as 'id แผนก hris',
    ja.`status_work` as 'id สถานะพนักงาน pis',
    (SELECT jp.name 
     FROM job_postion jp 
     WHERE jp.postion_id = ja.postion_id) AS 'ชื่อตำแหน่ง pis',
    (SELECT hp.name_th 
     FROM hris_positions hp 
     WHERE hp.id = ja.postion_id_new) AS 'ชื่อตำแหน่ง hris',
     CASE 
        WHEN ja.postion_id_new is NULL or ja.postion_id_new='' THEN (select pp2.position_id_new FROM job_postion as pp2 WHERE pp2.postion_id=ja.postion_id)
        ELSE ''
    END AS 'id ที่จะ map ให้ (เฉพาะคนที่ HRIS ไม่ได้ map ไว้)',
    (SELECT cc.node_name_th 
     FROM conf_com_const cc 
     WHERE cc.config_id = ja.position_level) AS 'ชื่อระดับตำแหน่ง pis',
    (SELECT cc2.node_name_th 
     FROM conf_com_const cc2 
     WHERE cc2.config_id = ja.position_level_new) AS 'ชื่อระดับตำแหน่ง hris',
    (select dp.dep_name FROM dep as dp WHERE dp.dep_id=ja.dep_id) as 'ชื่อฝ่าย pis',
    (select dph.name_th FROM hris_departments as dph WHERE dph.id=ja.dep_id_hris) as 'ชื่อฝ่าย hris',
    (select se.section_name FROM section as se WHERE se.section_id=ja.section_id) as 'ชื่อแผนก pis',
    (select seh.name_th FROM hris_departments as seh WHERE seh.id=ja.section_id_hris) as 'ชื่อแผนก hris',
    CASE 
        WHEN ja.status_work = 1 THEN 'พนักงาน'
        ELSE 'พ้นสภาพ'
    END AS 'สถานะพนักงาน pis',
    CASE 
        WHEN jp.postion_id IS NOT NULL AND jp.position_id_new = ja.postion_id_new THEN 'ตรง'
        ELSE 'ไม่ตรง'
    END AS 'map ตำแหน่งตรงกับ config',
    CASE 
        WHEN jp.postion_id IS NOT NULL AND jp.position_id_new <> ja.postion_id_new THEN jp.position_id_new
        ELSE NULL
    END AS 'id map ตำแหน่งที่ควรเป็น'
FROM 
    `job_applicant` ja
LEFT JOIN 
    `job_postion` jp 
    ON ja.postion_id = jp.postion_id
WHERE 
    1;




SELECT `applicant_id`,`th_name`,`th_surname`,`postion_id`,`dep_id`,`section_id`,`dep_id_hris`,`section_id_hris`,`status_work` FROM `job_applicant` WHERE 1




SELECT dep_id,dep_id_hris
FROM job_applicant
WHERE status_work=1
GROUP BY dep_id
HAVING COUNT(DISTINCT dep_id_hris) = 1;