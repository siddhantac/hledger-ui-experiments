class Transaction:
    date = ''
    txnid = ''
    desc = ''
    account1 = ''
    account2 = ''
    amount = 0

    def __init__(self, txnid, date, desc, account, amount):
        self.txnid = txnid
        self.date = date
        self.desc = desc
        self.account1 = account
        self.amount = amount
    def __str__(self):
        return 'id='+self.txnid+', date='+self.date+', desc='+self.desc+', account1='+self.account1+', account2='+self.account2+', amount='+str(self.amount)
