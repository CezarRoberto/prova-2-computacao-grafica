import cv2
import numpy as np

# Carrega a imagem
image = cv2.imread('prova 2\yami.jpg')

# Exibe a imagem original
cv2.imshow('Imagem Original', image)
cv2.waitKey(0)

# Realiza a mudança de escala (redimensionamento)
scaled_image = cv2.resize(image, None, fx=0.5, fy=0.5)

# Exibe a imagem redimensionada
cv2.imshow('Imagem Redimensionada', scaled_image)
cv2.waitKey(0)

# Define a matriz de translação
translation_matrix = np.float32([[1, 0, 50], [0, 1, 50]])

# Aplica a translação na imagem
translated_image = cv2.warpAffine(image, translation_matrix, (image.shape[1], image.shape[0]))

# Exibe a imagem com translação
cv2.imshow('Imagem com Translacao', translated_image)
cv2.waitKey(0)

# Realiza a rotação da imagem
rotation_angle = 90
rotation_matrix = cv2.getRotationMatrix2D((image.shape[1] / 2, image.shape[0] / 2), rotation_angle, 1)
rotated_image = cv2.warpAffine(image, rotation_matrix, (image.shape[1], image.shape[0]))

# Exibe a imagem rotacionada
cv2.imshow('Imagem Rotacionada', rotated_image)
cv2.waitKey(0)

# Fecha todas as janelas
cv2.destroyAllWindows()
