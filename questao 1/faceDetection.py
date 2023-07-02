import cv2


# Carrega o classificador pré-treinado para detecção de rosto
face_cascade = cv2.CascadeClassifier(cv2.data.haarcascades + 'haarcascade_frontalface_default.xml')


# Define o tamanho da janela
cv2.namedWindow('Detector de Rosto', cv2.WINDOW_NORMAL)
cv2.resizeWindow('Detector de Rosto', 800, 600)


# Inicia a captura de vídeo da webcam
cap = cv2.VideoCapture(0)


while True:
    # Lê um frame da captura de vídeo
    ret, frame = cap.read()


    # Converte o frame para escala de cinza
    gray = cv2.cvtColor(frame, cv2.COLOR_BGR2GRAY)


    # Detecta os rostos no frame
    rostos = face_cascade.detectMultiScale(gray, scaleFactor=1.1, minNeighbors=5, minSize=(60, 60))


    # Desenha retângulos ao redor dos rostos detectados
    for (x, y, w, h) in rostos:
        cv2.rectangle(frame, (x, y), (x+w, y+h), (0, 255, 0), 2)


    # Mostra o frame resultante
    cv2.imshow('Detector de Rosto', frame)


    # Sai do loop se a tecla 'Esc' for pressionada ou a janela for fechada
    if cv2.waitKey(1) == 27 or cv2.getWindowProperty('Detector de Rosto', cv2.WND_PROP_VISIBLE) < 1:
        break


# Libera os recursos
cap.release()
cv2.destroyAllWindows()
