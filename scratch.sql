select date_time_origination, calling_party_number, original_called_party_number, duration, orig_cause_value, 
orig_device_name, dest_device_name
from records
order by date_time_origination desc
limit 100;


-- Table: public.partition

-- DROP TABLE IF EXISTS public.partition;

CREATE TABLE IF NOT EXISTS public.partition
(
    "ID" character(3) COLLATE pg_catalog."default",
    "Partition_name" character(50) COLLATE pg_catalog."default"
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.partition
    OWNER to cdr;



SELECT DURATION,
	RECTAB.DATE_TIME_ORIGINATION,
	PARTTAB."ID" || RECTAB.CALLING_PARTY_NUMBER AS CALLING_NUMBER,
	RECTAB.CALLING_PARTY_NUMBER,
	PARTTAB."ID" AS PREF_CALLING,
	RECTAB.CALLING_PARTY_NUMBER_PARTITION, --calling_Party_Unicode_Login_User_ID,
 --original_Called_Party_Number_Partition,
 --original_Called_Party_Number,
 PARTTAB2."ID" || RECTAB.FINAL_CALLED_PARTY_NUMBER AS CALLED_NUMBER,
	RECTAB.FINAL_CALLED_PARTY_NUMBER,
	PARTTAB2."ID" AS PREF_CALLED,
	RECTAB.FINAL_CALLED_PARTY_NUMBER_PARTITION,
	RECTAB.DATE_TIME_CONNECT,
	RECTAB.DATE_TIME_DISCONNECT
FROM PUBLIC."records" AS RECTAB
LEFT OUTER JOIN PUBLIC."partition" AS PARTTAB ON RECTAB."calling_party_number_partition" = PARTTAB."Partition_name"
LEFT OUTER JOIN PUBLIC."partition" AS PARTTAB2 ON RECTAB."final_called_party_number_partition" = PARTTAB2."Partition_name" --    where calling_party_number = '414'

WHERE PARTTAB."ID" || RECTAB.CALLING_PARTY_NUMBER = '001460' -- отбор по исходящему номеру
 --WHERE PARTTAB2."ID" || RECTAB.FINAL_CALLED_PARTY_NUMBER = '001460' -- отбор по входящему номеру

	AND DURATION > 0 --    where final_Called_Party_Number = '220'

ORDER BY DATE_TIME_ORIGINATION DESC
LIMIT 1000;