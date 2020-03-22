from threading import Lock

class BankAccount:
    def __init__(self):
        self.amount = 0
        self.closed = True
        self.lock = Lock()

    def get_balance(self):
        if self.closed:
            raise ValueError("Your account was closed")

        return self.amount

    def open(self):
        if not self.closed:
            raise ValueError("You already have account")

        self.closed = False

    def deposit(self, amount):
        if self.closed:
            raise ValueError("Your account was closed")

        if amount < 0:
            raise ValueError("You can't deposite negative values")

        with self.lock:
            self.amount += amount

    def withdraw(self, amount):
        if self.closed:
            raise ValueError("Your account was closed")

        if amount < 0:
            raise ValueError("You can't deposit negative values")

        if self.amount < amount:
            raise ValueError("You can't withdraw more than your balance")

        with self.lock:
            self.amount -= amount

    def close(self):
        if self.closed:
            raise ValueError("You have already closed your account")            

        with self.lock:
            self.amount = 0
            self.closed = True
