


## Broadcast
### Input message example
		{
			"src": "n1",
			"dest": "c1",
			"body": {
				"type": "broadcast",
				"msg_id": 1,
			  "message": 1000
			}
		}
### Output message example
		{
			"src": "n1",
			"dest": "c1",
			"body": {
				"type": "broadcast_ok",
				"msg_id": 1,
				"in_reply_to": 1,
			  "message": 1000
			}
		}