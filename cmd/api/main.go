package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"max/pkg/database"

	"github.com/gorilla/mux"

	"max/internal/config"
	"max/internal/handler"
	"max/internal/middleware"
	"max/internal/repository"
	"max/internal/service"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Запуск банковского сервиса")

	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Подключение к базе данных
	db := database.NewDb(&cfg.Database)

	// Инициализация репозиториев
	userRepo := repository.NewUserRepository(db)
	accountRepo := repository.NewAccountRepository(db)
	cardRepo := repository.NewCardRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	creditRepo := repository.NewCreditRepository(db)
	// paymentRepo := repository.NewPaymentRepository(db)

	userService := service.NewUserService(userRepo, cfg.JWT.Secret)
	accountService := service.NewAccountService(accountRepo)
	cardService := service.NewCardService(cardRepo)
	transactionService := service.NewTransactionService(transactionRepo)
	creditService := service.NewCreditService(creditRepo)
	analyticsService := service.NewAnalyticsService(transactionRepo, creditRepo, accountRepo)

	// Инициализация обработчиков
	userHandler := handler.NewUserHandler(userService)
	accountHandler := handler.NewAccountHandler(accountService)
	cardHandler := handler.NewCardHandler(cardService)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	creditHandler := handler.NewCreditHandler(creditService)
	analyticsHandler := handler.NewAnalyticsHandler(analyticsService)

	// Инициализация маршрутизатора
	router := mux.NewRouter()

	// Middleware
	// router.Use(middleware.AuthMiddleware(cfg.JWT.Secret))

	// Публичные маршруты
	router.HandleFunc("/register", userHandler.Register).Methods("POST")
	router.HandleFunc("/login", userHandler.Login).Methods("POST")

	// Защищенные маршруты
	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware(cfg.JWT.Secret))

	// Маршруты для счетов
	api.HandleFunc("/accounts", accountHandler.CreateAccount).Methods("POST")
	api.HandleFunc("/accounts", accountHandler.GetAccounts).Methods("GET")
	api.HandleFunc("/accounts/{id}", accountHandler.GetAccount).Methods("GET")
	api.HandleFunc("/accounts/{id}/deposit", accountHandler.Deposit).Methods("POST")
	api.HandleFunc("/accounts/{id}/withdraw", accountHandler.Withdraw).Methods("POST")
	api.HandleFunc("/accounts/{id}/predict", accountHandler.PredictBalance).Methods("GET")

	// Маршруты для карт
	api.HandleFunc("/cards", cardHandler.CreateCard).Methods("POST")
	api.HandleFunc("/cards", cardHandler.GetCards).Methods("GET")
	api.HandleFunc("/cards/{id}", cardHandler.GetCard).Methods("GET")

	// Маршруты для транзакций
	api.HandleFunc("/transfer", transactionHandler.Transfer).Methods("POST")
	api.HandleFunc("/transactions", transactionHandler.GetTransactions).Methods("GET")

	// Маршруты для кредитов
	api.HandleFunc("/credits", creditHandler.CreateCredit).Methods("POST")
	api.HandleFunc("/credits", creditHandler.GetCredits).Methods("GET")
	api.HandleFunc("/credits/{id}", creditHandler.GetCredit).Methods("GET")
	api.HandleFunc("/credits/{id}/schedule", creditHandler.GetPaymentSchedule).Methods("GET")

	// Маршруты для аналитики

	api.HandleFunc("/analytics", analyticsHandler.GetAnalytics).Methods("GET")
	api.HandleFunc("/analytics/credit-load", analyticsHandler.GetCreditLoad).Methods("GET")

	// Запуск HTTP-сервера
	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Запуск сервера в горутине
	go func() {
		log.Infof("Сервер запущен на порту %s", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка запуска сервера: %v", err)
		}
	}()

	// Запуск шедулера для обработки просроченных платежей
	// go creditService.StartPaymentScheduler(12 * time.Hour)

	// Обработка сигналов для корректного завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Завершение работы сервера...")

	// Создаем контекст с таймаутом для корректного завершения
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка при завершении сервера: %v", err)
	}

	log.Info("Сервер успешно остановлен")
}
