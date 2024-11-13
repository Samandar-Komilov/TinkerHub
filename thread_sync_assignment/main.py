import threading
import time
from queue import Queue
from random import randint

MAX_ROOMS = 3
available_rooms = MAX_ROOMS

room_mutex = threading.Lock()
cv = threading.Condition(room_mutex)
waiting_queue = Queue()

def reserve_room(student_id):
    global available_rooms

    with room_mutex:
        waiting_queue.put(student_id)
        print(f"Student {student_id} is waiting to reserve a room.")

        cv.wait_for(lambda: waiting_queue.queue[0] == student_id and available_rooms > 0)

        waiting_queue.get()
        available_rooms -= 1
        print(f"Student {student_id} reserved a room. Rooms available: {available_rooms}")

    study_time = randint(1,5)
    time.sleep(study_time)

    with room_mutex:
        available_rooms += 1
        print(f"Student {student_id} released a room in {study_time} seconds. Rooms available: {available_rooms}")
        cv.notify_all()

def main():
    students = []
    for i in range(10):
        t = threading.Thread(target=reserve_room, args=(i+1,))
        students.append(t)
        t.start()

    for t in students:
        t.join()

if __name__ == "__main__":
    main()
