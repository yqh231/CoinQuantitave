import smtplib
from email.mime.text import MIMEText
from email.header import Header

from config import config


def send_email(subject, text):
    message = MIMEText(text, 'plain', 'utf-8')
    message['From'] = Header(config.email, 'utf-8')
    message['To'] = Header(config.email, 'utf-8')
    message['Subject'] = Header(subject, 'utf-8')

    smtp = smtplib.SMTP() 
    smtp.connect('smtp.163.com') 
    smtp.login(config.email, config.pass_word) 
    smtp.sendmail(config.email, [config.email], message.as_string()) 
    smtp.quit()
