package console

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"portpower/internal/audit"
	"portpower/internal/berth"
	"portpower/internal/cable"
	"portpower/internal/disconn"
	"portpower/internal/freq"
	"portpower/internal/meter"
	"portpower/internal/ns"
	"portpower/internal/phase"
	"portpower/internal/quota"
	"portpower/internal/ship"
	"portpower/internal/store"
)

type Runtime struct {
	harbor  *ns.Harbor
	catalog *ns.Catalog
	store   *store.Store
	ledger  *audit.Ledger
	quotas  *quota.Manager
	ships   *ship.Registry
	fleet   *ship.Fleet
	berths  *berth.Mapping
	berthReg *berth.Registry
	binder  *meter.Binder
	freq    *freq.Syncer
	sensor  *cable.Sensor
	filter  *cable.Filter
	cable   *cable.Cable
	cables  *cable.Inventory
	tariff  *meter.Tariff
	monitor *cable.Monitor
	sched   *berth.Scheduler
	history *ship.History
	phase   *phase.State
	shore   *disconn.Breaker
	protect *disconn.Protector
	limit   *disconn.LimitEnforcer
}

type Server struct {
	runtime *Runtime
	router  *chi.Mux
}

func newRuntime(harbor *ns.Harbor, st *store.Store) *Runtime {
	filter := cable.NewFilter(3)
	monitorFilter := cable.NewFilter(5)
	rt := &Runtime{
		harbor:  harbor,
		catalog: ns.NewCatalog("A1", "A2", "B1", "B2", "C1"),
		store:   st,
		ledger:  audit.NewLedger(st),
		quotas:  quota.NewManager(st),
		ships:   ship.NewRegistry(st),
		fleet:   ship.NewFleet(),
		berths:  berth.NewMapping(st),
		berthReg: berth.NewRegistry(),
		binder:  meter.NewBinder(st),
		freq:    freq.NewSyncer(st, 50),
		sensor:  cable.NewSensor(25),
		filter:  filter,
		cable:   cable.NewCable("shore-1", 90),
		cables:  cable.NewInventory(),
		tariff:  meter.NewTariff(0.95),
		monitor: cable.NewMonitor(monitorFilter),
		sched:   berth.NewScheduler(st),
		history: ship.NewHistory(st),
		phase:   phase.NewState(phase.SequenceForward),
		shore:   disconn.NewBreaker("shore-main"),
		protect: disconn.NewProtector(90, filter),
		limit:   disconn.NewLimitEnforcer(ship.NewRegistry(st)),
	}
	return rt
}

func NewServer(harbor *ns.Harbor, st *store.Store) *Server {
	s := &Server{runtime: newRuntime(harbor, st), router: chi.NewRouter()}
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.routes()
	return s
}

func (s *Server) Handler() http.Handler {
	return s.router
}

func (s *Server) routes() {
	s.router.Get("/health", s.handleHealth)
	s.router.Post("/connect", s.handleConnect)
	s.router.Post("/disconnect", s.handleDisconnect)
	s.router.Post("/meter", s.handleMeter)
	s.router.Get("/usage", s.handleUsage)
	s.router.Get("/berths", s.handleBerths)
	s.router.Post("/retrofit", s.handleRetrofit)
	s.router.Get("/berth/status", s.handleBerthStatus)
	s.router.Get("/ships", s.handleShips)
	s.router.Get("/audit", s.handleAudit)
	s.router.Get("/audit/latest", s.handleAuditLatest)
	s.router.Get("/quota", s.handleQuotaList)
	s.router.Post("/quota", s.handleQuotaSet)
	s.router.Post("/snapshot", s.handleSnapshot)
	s.router.Post("/restore", s.handleRestore)
	s.router.Get("/freq", s.handleFreqGet)
	s.router.Post("/freq", s.handleFreqSet)
	s.router.Get("/cables", s.handleCables)
	s.router.Post("/allocate", s.handleAllocate)
	s.router.Post("/vacate", s.handleVacate)
	s.router.Get("/billing", s.handleBilling)
	s.router.Post("/quota/reserve", s.handleQuotaReserve)
	s.router.Post("/quota/release", s.handleQuotaRelease)
	s.router.Get("/audit/summary", s.handleAuditSummary)
	s.router.Post("/breaker/trip", s.handleBreakerTrip)
	s.router.Post("/breaker/reset", s.handleBreakerReset)
	s.router.Post("/reservation", s.handleReservationBook)
	s.router.Get("/reservations", s.handleReservationList)
	s.router.Delete("/reservation", s.handleReservationCancel)
	s.router.Get("/ship/history", s.handleShipHistory)
	s.router.Get("/meter/period", s.handleMeterPeriod)
	s.router.Get("/cable/monitor", s.handleCableMonitor)
	s.router.Get("/breaker/interlock", s.handleBreakerInterlock)
	s.router.Get("/status", s.handleStatus)
	s.router.Get("/freq/samples", s.handleFreqSamples)
}
