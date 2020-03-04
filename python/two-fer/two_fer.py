def two_fer(name=""):
    msg = "One for {}, one for me."
    if name == "":
        name = "you"

    return msg.format(name)
