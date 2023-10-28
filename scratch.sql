select date_time_origination, calling_party_number, original_called_party_number, duration, orig_cause_value, 
orig_device_name, dest_device_name
from records
order by date_time_origination;