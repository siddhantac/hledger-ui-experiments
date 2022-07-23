import csv
import subprocess
from transaction import Transaction

# result=subprocess.run(
#         ["hledger", "reg", "-p", "2022/01/01 to 2022/01/03", "-O", "csv"],
#         check=True,
#         capture_output=True,
#         )
#
# lines=result.stdout.splitlines()
#
transactions = {}
#
# for l in lines:
#     tx = Transaction()
#     items = l.split(b',')
#     tx.txnid = str(items[0])
#     transactions[tx.txnid] = tx
#     print(items)
#     print(tx)

with open('data.csv') as csvfile:
    reader = csv.reader(csvfile)
    for row in reader:
        tx = Transaction(
                txnid=row[0], 
                date=row[1],
                desc=row[3],
                account=row[4],
                amount=row[5],
                )
        tx2 = transactions.get(tx.txnid, None)
        if tx2 is None:
            transactions[tx.txnid] = tx
        else:
            tx2.account2 = row[4]
        # print(tx2)
        print(' '.join(row))
